const pad2 = (n: number) => n.toString().padStart(2, '0')
export default class Clock {
    private readonly date: Date

    constructor(hour: number, minute: number = 0) {
        this.date = new Date(2000, 0, 0, hour, minute)
    }

    private getHour = () => this.date.getHours()
    private getMinute = () => this.date.getMinutes()

    plus = (minute: number) => new Clock(this.getHour(), this.getMinute() + minute)
    minus = (minute: number) => new Clock(this.getHour(), this.getMinute() - minute)
    equals = (otherClock: Clock) => this.getHour() === otherClock.getHour() && this.getMinute() === otherClock.getMinute()
    toString = () => pad2(this.getHour()) + ':' + pad2(this.getMinute())
}