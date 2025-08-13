package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"github.com/gorilla/mux"
	qrcode "github.com/skip2/go-qrcode"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"time"
)

type BarcodeSimples struct {
	Code     string
	Datahora string
}

type Barcodes struct {
	Barcodes []BarcodeSimples
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func write(filename string, bs Barcodes) {
	data, err := json.MarshalIndent(bs, "", "  ")
	if err != nil {
		log.Fatal("Erro ao criar JSON:", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatal("Erro ao criar arquivo:", err)
	}

	fmt.Println("Arquivo criado com sucesso:", filename)
}

func writeToFileSimples(barcode string) {
	filename := "simples.json"

	if !fileExists(filename) {
		fmt.Println("Arquivo não existe, criando com dados padrão...")

		bs := Barcodes{
			Barcodes: []BarcodeSimples{
				{Code: barcode, Datahora: time.Now().Format("02/01/2006 15:04")},
			},
		}
		write(filename, bs)
	}

	// Agora lê o arquivo
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Erro ao abrir arquivo:", err)
	}
	defer file.Close()

	var bs Barcodes
	err = json.NewDecoder(file).Decode(&bs)
	if err != nil {
		log.Fatal("Erro ao decodificar JSON:", err)
	}

	bs.Barcodes = append(bs.Barcodes, BarcodeSimples{Code: barcode, Datahora: time.Now().Format("02/01/2006 15:04")})
	write(filename, bs)

}

func resetSimples() {
	filename := "simples.json"

	if !fileExists(filename) {
		return
	}
	write(filename, Barcodes{Barcodes: []BarcodeSimples{}})
}

func generateSelfSignedCert() (tls.Certificate, error) {
	// Gera chave privada RSA de 2048 bits
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return tls.Certificate{}, err
	}

	// Certificado válido por 1 ano
	serialNumber, err := rand.Int(rand.Reader, big.NewInt(1<<62))
	if err != nil {
		return tls.Certificate{}, err
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Meu Servidor Go"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(365 * 24 * time.Hour), // 1 ano
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IPAddresses:           []net.IP{net.ParseIP("127.0.0.1")}, // válido apenas para localhost
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return tls.Certificate{}, err
	}

	// Codifica chave privada e certificado em PEM
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})

	// Cria certificado TLS
	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		return tls.Certificate{}, err
	}

	return cert, nil
}

func generate_qrcode() {
	url, err := getIPv4()
	if err != nil {
		log.Fatal(err)
	}
	qr, err := qrcode.New("https://"+url+":8443", qrcode.Medium)
	if err != nil {
		panic(err)
	}
	// false = borda, true = sem borda
	fmt.Println(qr.ToSmallString(false))
}

func getIPv4() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		}

		// Pega apenas IPv4 que não seja loopback
		if ip != nil && ip.To4() != nil && !ip.IsLoopback() {
			return ip.String(), nil
		}
	}
	return "", fmt.Errorf("nenhum IPv4 encontrado")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/simples/reset", func(w http.ResponseWriter, r *http.Request) {
		resetSimples()
		w.Write([]byte("Arquivo resetado com sucesso"))
	}).Methods("GET")

	r.HandleFunc("/simples/all", func(w http.ResponseWriter, r *http.Request) {
		filename := "simples.json"
		data, err := os.ReadFile(filename) 
		if err != nil {
			log.Fatal("Erro ao abrir arquivo:", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}).Methods("GET")

	r.HandleFunc("/simples/{barcode}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		barcode := vars["barcode"]
		writeToFileSimples(barcode)
		w.Write([]byte("Barcode adicionado com sucesso"))
	}).Methods("GET")

	// Gera certificado TLS em memória
	cert, err := generateSelfSignedCert()
	if err != nil {
		log.Fatal("Erro ao gerar certificado:", err)
	}

	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}}

	server := &http.Server{
		Addr:      ":8443",
		Handler:   r,
		TLSConfig: tlsConfig,
	}

	log.Println("Servidor HTTPS rodando em https://localhost:8443")
	generate_qrcode()
	log.Fatal(server.ListenAndServeTLS("", "")) // certificados vêm de tlsConfig, então os caminhos ficam vazios
}
