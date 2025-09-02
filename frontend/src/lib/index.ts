// place files you want to import through the `$lib` alias in this folder.
export const FILE_PATH = "file:///Users/juniper/Documents/coding/archeyeve/images/0/1.jpeg"

export type Observation = {
    id: number
    userid: number
    name: string
    address: string
    styles: string[]
    type: string[]
    condition: string[]
    lng: number
    lat: number
    year: number
    imgcount: number
}


export type Style = {
    id: number
    name: string
    era: string
}