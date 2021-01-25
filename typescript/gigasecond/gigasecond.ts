class Gigasecond {
  constructor(private inputDate: Date) {
  }
  date() {
    return new Date(this.inputDate.getTime() + 1000000000 * 1000)
  }
}

export default Gigasecond
