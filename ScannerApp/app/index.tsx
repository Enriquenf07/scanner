import useApi, { createInstance } from "@/network";
import AsyncStorage from "@react-native-async-storage/async-storage";
import { useRouter } from "expo-router";
import { Button, Text, TouchableOpacity, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";

export default function Index() {
  const router = useRouter();
  const [url, loading] = useApi()

  const reset = async () => {
    if (loading) return
    const api = createInstance(url)
    try {
      api.get(`/simples/reset`)
    } catch (e) {
      console.log(e)
    }
  }
  return (
    <SafeAreaView style={{
      flex: 1,
      justifyContent: 'center',
      marginHorizontal: 16,
    }}>
      <View style={{ gap: 10 }}>
        <Button title='scanner' onPress={() => router.navigate('/scanner')}>
        </Button>
        <Button title='reset' onPress={() => reset()}>
        </Button>
      </View>
    </SafeAreaView>
  )
}

