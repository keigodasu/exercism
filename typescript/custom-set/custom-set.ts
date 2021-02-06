export default class CustomSet {

    constructor(private set: number[] = []) {
    }

    empty = () => this.set.length === 0
    contains = (n: number) => this.set.includes(n)
    subset = (other: CustomSet) => this.set.every(other.contains)
    disjoint = (other: CustomSet) => this.set.every(e => !other.contains(e))
    eql = (other: CustomSet) => this.set.length === other.set.length && this.subset(other)
    add = (n: number) => {
        if (!this.contains(n)) {
            this.set.push(n)
        }
        return this
    }
    intersection = (other: CustomSet) => new CustomSet(this.set.filter(other.contains))
    difference = (other: CustomSet) => new CustomSet(this.set.filter(el => !other.contains(el)))
    union = (other: CustomSet) => {
        this.set.forEach(e => other.add(e))
        return other
    }
}