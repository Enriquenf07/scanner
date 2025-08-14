import useApi, { createInstance } from "@/network";
import AsyncStorage from "@react-native-async-storage/async-storage";
import { useRouter } from "expo-router";
import { Button, Text, TouchableOpacity, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";

export default function Index() {
  const router = useRouter();
  return (
    <SafeAreaView style={{
      flex: 1,
      justifyContent: 'center',
      marginHorizontal: 16,
    }}>
      <View style={{ gap: 10 }}>
        <Button title='scanner' onPress={() => router.navigate('/scanner')}>
        </Button>
      </View>
    </SafeAreaView>
  )
}

