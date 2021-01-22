class Transcriptor {
    private mapping: {[name:string]:string} = {
        'G': 'C',
        'C': 'G',
        'T': 'A',
        'A': 'U',
    }

    private error() {
        throw Error('Invalid input DNA.')
    }

    toRna(dna: string) {
        return dna.split('').map((el) => this.mapping[el] || this.error()).join('')
    }
}

export default Transcriptor
