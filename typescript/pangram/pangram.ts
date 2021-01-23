class Pangram {
    constructor(private input: string) {
    }

    isPangram(): boolean {
        if ((this.input.toLowerCase().match(/([a-z])(?!.*\1)/g)||[]).length === 26) {
            return true
        }
        return false
    }
}

export default Pangram
