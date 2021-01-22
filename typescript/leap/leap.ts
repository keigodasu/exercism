function isLeapYear(year: number): boolean {
    return new Date(year, 2, 0).getDate() == 29
}

export default isLeapYear
