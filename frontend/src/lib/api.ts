import type { Observation } from "$lib";

export async function fetchObservation(id: number): Promise<Observation> {
    return {
        id: id,
        userid: 0,
        name: 'Igreja de São João Baptista',
        lng: 0.0,
        lat: 0.0,
        style: ["Manueline", "Baroque"],
        year: 2025,
        imgcount: 1
    }
}

const observations: Observation[] = [{
    id: 0,
    userid: 0,
    name: "Praça de Touros do Campo Pequeno",
    lng: -9.145202,
    lat: 38.742590,
    style: ["neo-mudejar"],
    year: 1892,
    imgcount: 1 
}, {
    id: 1,
    userid: 0,
    name: "Torre do Tombo",
    lng: -9.156460,
    lat: 38.754630,
    style: ["post-modernism","brutalism"],
    year: 1990,
    imgcount: 1
}, {
    id: 2,
    userid: 0,
    name: "Torre da Praça do Areeiro",
    lng: -9.133270,
    lat: 38.742910,
    style: ["português suave"],
    year: 1938,
    imgcount: 1
}, {
    id: 3,
    userid: 0,
    name: "St George's Church",
    lng: -9.160482,
    lat: 38.716496,
    style: ["neo-romanesque"],
    year: 1889,
    imgcount: 1 
}, {
    id: 4,
    userid: 0,
    name: "EDP Sede II",
    lng: -9.148990,
    lat: 38.707162,
    style: ["brutalism"],
    year: 2024,
    imgcount: 1 
}, {
    id: 5,
    userid: 0,
    name: "Igreja de São João Baptista",
    lng: -8.414697,
    lat: 39.603648,
    style: ["manueline"],
    year: 1500,
    imgcount: 1
}, {
    id: 6,
    userid: 0,
    name: "Sé de Braga",
    lng: -8.427396,
    lat: 41.549916,
    style: ["romanesque", "manueline", "baroque"],
    year: 1089,
    imgcount: 1
}, {
    id: 7,
    userid: 0,
    name: "Santuário do Monte de Santa Luzia",
    lng: -8.835104,
    lat: 41.701517,
    style: ["neo-byzantine", "neo-romanesque", "neo-gothic"],
    year: 1904,
    imgcount: 1
}, {
    id: 8,
    userid: 0,
    name: "Citânia de Briteiros",
    lng: -8.315402,
    lat: 41.527687,
    style: ["castro", "iron-age"],
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

export type Search = Partial<BaseSearch & (CircleSearch | AreaSearch)>

export async function doSearchREAL(s: Search): Promise<Observation[]> {
    Object.keys(s).forEach(k => {
        // @ts-ignore
        s[k] = s[k].toString()
    })
    const searchParams = new URLSearchParams(s as Record<string, string>)
    const url = "http://localhost:8080/search?" + searchParams.toString()
    let response = await fetch(url)

    return response.json()
}

export async function doSearch(s: Search): Promise<Observation[]> {
    let result: Observation[] = []
    observations.forEach((o) => {
        if (s.style == o.style[0] || s.style == o.style[1] || s.style == o.style[2])
        result.push(o)
    })

    return result
}

