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

export function getMimeType(file: any, fallback = null) {
  const byteArray = new Uint8Array(file).subarray(0, 4)
  let header = ''

  for (let i = 0; i < byteArray.length; i++) {
    header += byteArray[i].toString(16)
  }

  switch (header) {
    case '89504e47':
      return 'image/png'
    case 'ffd8ffe0':
    case 'ffd8ffe1':
    case 'ffd8ffe2':
    case 'ffd8ffe3':
    case 'ffd8ffe8':
      return 'image/jpeg'
    default:
      return fallback
  }
}
