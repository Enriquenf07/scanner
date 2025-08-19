import { useQuery } from "@tanstack/react-query"

interface Item {
    id: number,
    datahora: string,
    produto: string,
    code: string
}

function Itens() {
    const { isPending, error, data } = useQuery({
        queryKey: ['itens'],
        queryFn: () =>
            fetch('/scanner/barcode').then((res) =>
                res.json(),
            ),
    })

    if (isPending) return 'Carregando...'

    if (error) return 'Ocorreu um erro ' + error.message

    return data && data.map((item: Item) => (
        <div key={item.id}>
            <p>{item.datahora}</p>
            <p>{item.code}</p>
            <p>{item.produto}</p>
        </div>
    ))
}

function Excel() {
    const { isPending, error, refetch } = useQuery({
        queryKey: ['excel'],
        queryFn: () =>
            fetch('/scanner/excel') 
                .then(res => {
                    if (!res.ok) throw new Error('Erro ao gerar Excel');
                    return true; 
                }),
        enabled: false,

    })

    if (error) return (
        <div>{error?.message}</div>
    )

    return (
        <div className="flex gap-2">
            <button onClick={() => refetch()}>Atualizar</button>
            {!isPending && <button>Download</button>}
        </div>
    )
}

export default function Home() {
    return (
        <div className="bg-zinc-200 w-full min-h-screen">
            <Itens />
            <Excel />
        </div>
    )
}