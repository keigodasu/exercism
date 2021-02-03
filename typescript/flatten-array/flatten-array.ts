export default class FlattenArray {
    static flatten(arr: any): number[] {
        return arr.reduce((acc: any, e:any): number[] => {
            if (Array.isArray(e)) return [...acc, ...this.flatten(e)]
            if (typeof e === "number") acc.push(e)
            return acc
        }, [])
    }
}
