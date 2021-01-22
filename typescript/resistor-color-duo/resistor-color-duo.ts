export class ResistorColor {
  readonly colors: string[]
  private registerValues: { [name: string]: number } = {
    black: 0,
    brown: 1,
    red: 2,
    orange: 3,
    yellow: 4,
    green: 5,
    blue: 6,
    violet: 7,
    grey: 8,
    white: 9
  }

  constructor(colors: string[]) {
    if (colors.length < 2) throw Error('At least two colors need to be present')
    this.colors = colors
  }
  value = (): number => {
    return this.registerValues[this.colors[0]] * 10 + this.registerValues[this.colors[1]]
  }
}
