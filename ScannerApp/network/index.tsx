import AsyncStorage from "@react-native-async-storage/async-storage";
import axios, { AxiosInstance } from "axios";
import { useEffect, useState } from "react";

export const createInstance = (url: string) => axios.create({
  baseURL: `${url}:8080`,
  timeout: 1000,
});

export default function useApi(): [string, boolean] {
  const [loading, setLoading] = useState<boolean>(false)
  const [url, setUrl] = useState<string>('')

  const read = async () => {
    setLoading(true)
    try {
      const value = await AsyncStorage.getItem('url');
      if (value !== null) {
        setUrl(value)
      }
    } catch (e) {
      console.log(e)
    }finally{
      setLoading(false)
    }
  }
  useEffect(() => {
    read()
  }, [])
  
  return [url, loading]

}