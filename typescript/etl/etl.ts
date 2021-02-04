function transform(input: {[key: string]: string[]}):{[key: string]: number} {
    const result: {[key: string]: number} = {}
    for (const score in input) {
        for(const letter of input[score])  {
            result[letter.toLowerCase()] = Number(score)
        }
    }

    return result
}

export default transform