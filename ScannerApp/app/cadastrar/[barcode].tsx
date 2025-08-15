import { api } from "@/network"
import { useRoute } from "@react-navigation/native"
import axios from "axios"
import { useRouteInfo, useRouter } from "expo-router/build/hooks"
import { useState } from "react"
import { Button, Text, TextInput, View } from "react-native"
import { SafeAreaView } from "react-native-safe-area-context"

export default function Cadastrar() {
    const route = useRouteInfo()
    const router = useRouter()
    const [nome, setNome] = useState<string>('')
    const salvar = async() => {
        try {
            console.log({ produto: nome, barcode: route.params.barcode })
            const response = await axios.post('https://v2202508293879372462.ultrasrv.de/scanner/barcode', { produto: nome, barcode: route.params.barcode })
            console.log(response)
            router.navigate('/')
        } catch (e: any) {
            console.log(e.message)
        }
    }
    return (
        <SafeAreaView>
            <View>
                <TextInput
                    multiline
                    numberOfLines={4}
                    maxLength={40}
                    onChangeText={text => setNome(text)}
                    value={nome}
                />
                <Button title='salvar' onPress={salvar} />
            </View>
        </SafeAreaView>
    )
}