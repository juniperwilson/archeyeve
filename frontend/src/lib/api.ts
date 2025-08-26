import type { Observation, Style } from "$lib";

const observations: Observation[] = [{
    id: 0,
    userid: 0,
    name: "Praça de Touros do Campo Pequeno",
    address: "",
    lng: -9.145202,
    lat: 38.742590,
    styles: ["neo-mudejar"],
    year: 1892,
    imgcount: 1 
}, {
    id: 1,
    userid: 0,
    name: "Torre do Tombo",
    address: "",
    lng: -9.156460,
    lat: 38.754630,
    styles: ["post-modernism","brutalism"],
    year: 1990,
    imgcount: 1
}, {
    id: 2,
    userid: 0,
    name: "Torre da Praça do Areeiro",
    address: "",
    lng: -9.133270,
    lat: 38.742910,
    styles: ["português suave"],
    year: 1938,
    imgcount: 1
}, {
    id: 3,
    userid: 0,
    name: "St George's Church",
    address: "",
    lng: -9.160482,
    lat: 38.716496,
    styles: ["neo-romanesque"],
    year: 1889,
    imgcount: 1 
}, {
    id: 4,
    userid: 0,
    name: "EDP Sede II",
    address: "",
    lng: -9.148990,
    lat: 38.707162,
    styles: ["brutalism"],
    year: 2024,
    imgcount: 1 
}, {
    id: 5,
    userid: 0,
    name: "Igreja de São João Baptista",
    address: "",
    lng: -8.414697,
    lat: 39.603648,
    styles: ["manueline"],
    year: 1500,
    imgcount: 1
}, {
    id: 6,
    userid: 0,
    name: "Sé de Braga",
    address: "",
    lng: -8.427396,
    lat: 41.549916,
    styles: ["romanesque", "manueline", "baroque"],
    year: 1089,
    imgcount: 1
}, {
    id: 7,
    userid: 0,
    name: "Santuário do Monte de Santa Luzia",
    address: "",
    lng: -8.835104,
    lat: 41.701517,
    styles: ["neo-byzantine", "neo-romanesque", "neo-gothic"],
    year: 1904,
    imgcount: 1
}, {
    id: 8,
    userid: 0,
    name: "Citânia de Briteiros",
    address: "",
    lng: -8.315402,
    lat: 41.527687,
    styles: ["castro", "iron-age"],
    year: -200,
    imgcount: 1
}]

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
    // let s: Style = {name: name}
    // let searchparams = new URLSearchParams(s as Record<string, string>)
    const url = "http://localhost:8080/style?name=" + name
    let response = await fetch(url)

    return response.json()
}

