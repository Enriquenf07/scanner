
import { BarcodeScanningResult, CameraType, CameraView, useCameraPermissions } from "expo-camera";
import { Link, useRouter } from "expo-router";
import { StatusBar } from "expo-status-bar";
import { useState } from "react";
import { Button, StyleSheet, Text, TouchableOpacity, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";


export default function Scanner({ onScan }: { onScan: (result: BarcodeScanningResult) => void; }) {
    const [permission, requestPermission] = useCameraPermissions();
    const [dataCache, setDataCache] = useState<string[]>([])
    const router = useRouter()

    if (!permission) {
        return <View />;
    }

    if (!permission.granted) {
        return (
            <View style={styles.container}>
                <Text style={styles.message}>We need your permission to show the camera</Text>
                <Button onPress={requestPermission} title="grant permission" />
            </View>
        );
    }



    const onScanWrapper = (result: BarcodeScanningResult) => {
        if (dataCache.includes(result.data)) {
            return
        }
        setDataCache(prev => {
            if (!prev.includes(result.data)) {
                return [...prev, result.data];
            }
            return prev
        });
        onScan(result)
    }



    return (
        <SafeAreaView style={styles.container}>
            <StatusBar hidden />
            <CameraView
                style={styles.camera}
                facing={"back"}
                onBarcodeScanned={onScanWrapper}
            />
            <View>
                <Text>{dataCache.findLast(predicate => predicate)}</Text>
                <Link href={`/cadastrar/${dataCache.findLast(predicate => predicate)}`}>Continuar</Link>
            </View>
        </SafeAreaView>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        justifyContent: 'center',
    },
    message: {
        textAlign: 'center',
        paddingBottom: 10,
    },
    camera: {
        flex: 1,
    },
    buttonContainer: {
        flex: 1,
        flexDirection: 'row',
        backgroundColor: 'transparent',
        margin: 64,
    },
    button: {
        flex: 1,
        alignSelf: 'flex-end',
        alignItems: 'center',
    },
    text: {
        fontSize: 24,
        fontWeight: 'bold',
        color: 'white',
    },
});