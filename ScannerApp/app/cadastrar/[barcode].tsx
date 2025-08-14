import { api } from "@/network"
import { useRoute } from "@react-navigation/native"
import { useRouteInfo } from "expo-router/build/hooks"
import { useState } from "react"
import { Button, TextInput, View } from "react-native"
import { SafeAreaView } from "react-native-safe-area-context"

export default function Cadastrar() {
    const route = useRouteInfo()
    
    const [nome, setNome] = useState<string>('')
    const salvar = () => {
        console.log({produto: nome, barcode: route.params.barcode})
        api.post('/barcode/', {produto: nome, barcode: route.params.barcode})
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