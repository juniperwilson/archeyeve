// place files you want to import through the `$lib` alias in this folder.
//export const FILE_PATH = "file:///Users/juniper/documents/coding/archeyeve/images/"
export const FILE_PATH = "/"
export type Observation = {
    id: number
    userid: number
    name: string
    lng: number
    lat: number
    style: string[]
    year: number
    imgcount: number
}