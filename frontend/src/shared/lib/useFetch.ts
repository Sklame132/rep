import { useState, useEffect, useMemo } from "react";

export type FetchState<T> = {
    loading: boolean;
    data: T | undefined;
    error: string | undefined;
}

const API_URL = "http://localhost:5050/api/v1"

const DEFAULT_OPTIONS = {
    headers: {"Content-Type": "application/json"},
}

export function useFetch<T>(url: string, options?: any): FetchState<T> {
    const [status, setStatus] = useState<FetchState<T>>({
        loading: false,
        data: undefined,
        error: undefined,
    })

    useEffect(() => {
        setStatus({ loading: true, data: undefined, error: undefined });
        const fullURL = API_URL + url;

        fetch(fullURL, { ...DEFAULT_OPTIONS, ...options })
        .then((res) => {
            if (!res.ok) {
                throw new Error(`${res.status}`)
            }
            return res.json();
        })
        .then((data: T) => {
            setStatus({ loading: false, data: data, error: undefined})
        })
        .catch((error) => {
            setStatus({loading: false, data: undefined, error: error.message})
        });
    }, [url]);

    return useMemo(() => ({ ...status}), [status]);
}