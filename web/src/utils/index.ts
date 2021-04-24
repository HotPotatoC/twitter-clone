// source: https://stackoverflow.com/a/56934067/12472142
export function getDaysInMonth(month: number, year: number): Date[] {
  var date = new Date(year, month, 1)
  var days = []
  while (date.getMonth() === month) {
    days.push(new Date(date))
    date.setDate(date.getDate() + 1)
  }
  return days
}

export function getYears() {
  const currentYear = new Date().getFullYear()
  return range(currentYear, 1900, -1)
}

export function range(start: number, stop: number, step: number) {
  return Array.from(
    { length: (stop - start) / step + 1 },
    (_, i) => start + i * step
  )
}
