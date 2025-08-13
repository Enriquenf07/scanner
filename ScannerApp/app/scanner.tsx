import Scanner from "@/components/Scanner";
import useApi, { createInstance } from "@/network";

import { BarcodeScanningResult } from "expo-camera";
import { SafeAreaView, Text } from "react-native";

export default function ScannerPage() {
    const onScan = (result: BarcodeScanningResult) => {
    }
    return (
        <Scanner onScan={onScan} />
    )
}