import type { Observation, Style } from "$lib";

export type BaseSearch = {
    style: string
    befyear: number
    aftyear: number
}

export type CircleSearch = {
    lat: number
    lng: number
    rad: number
}

export type AreaSearch = {
    lat1: number
    lng1: number
    lat2: number
    lng2: number
}

export type ObservationID = {
    id: string
}

export type Search = Partial<BaseSearch & (CircleSearch | AreaSearch)>

export async function doSearch(s: Search): Promise<Observation[]> {
    Object.keys(s).forEach(k => {
        // @ts-ignore
        s[k] = s[k].toString()
    })
    const searchParams = new URLSearchParams(s as Record<string, string>)
    const url = "http://localhost:8080/search?" + searchParams.toString()
    let response = await fetch(url)

    return response.json()
}

export async function getObservation(id: number): Promise<Observation> {

    let s: ObservationID = {id: id.toString()}
    let searchparams = new URLSearchParams(s as Record<string, string>)
    const url = "http://localhost:8080/observation?" + searchparams
    let response = await fetch(url)

    return response.json()
}

export async function fetchStyle(name: string): Promise<Style> {
    const url = "http://localhost:8080/style?name=" + name
    let response = await fetch(url)

    return response.json()
}

export async function getBuildings(): Promise<Observation[]> {
    return doSearch({})
}

export async function getStyles(): Promise<Style[]> {
    const url = "http://localhost:8080/style?"
    let response = await fetch(url)

    return response.json()
}

