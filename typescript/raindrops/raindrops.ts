class Raindrops {

    convert(num:number) :string{
        let out = ''
        if (num % 3 == 0) {out+="Pling"}
        if (num % 5 == 0) {out+="Plang"}
        if (num % 7 == 0) {out+="Plong"}
        if (out === '') {out = num.toString()}

        return out
    }
}

export default Raindrops
