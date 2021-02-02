class SpaceAge {
    private earthYearBySecond = 31557600;

    constructor(public seconds: number) {
    }

    roundedEarthYear(num: number): number {
        return Math.round(num * 100 ) / 100
    }

    onEarth():number {
        return this.roundedEarthYear(this.seconds /this.earthYearBySecond)
    }

    onMercury():number {
        return this.roundedEarthYear(this.seconds/this.earthYearBySecond/0.2408467)
    }

    onVenus():number {
        return this.roundedEarthYear(this.seconds/this.earthYearBySecond/0.61519726)
    }

    onMars():number {
        return this.roundedEarthYear(this.seconds/this.earthYearBySecond/1.8808158)
    }

    onJupiter():number {
        return this.roundedEarthYear(this.seconds/this.earthYearBySecond/11.862615)
    }

    onSaturn(): number {
        return this.roundedEarthYear(this.seconds/this.earthYearBySecond/29.447498)
    }

    onUranus(): number {
        return this.roundedEarthYear(this.seconds/this.earthYearBySecond/84.016846)
    }

    onNeptune(): number {
        return this.roundedEarthYear(this.seconds/this.earthYearBySecond/164.79132)
    }
}


export default SpaceAge