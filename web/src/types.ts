export type ResponseItems<T> = {
  items: T[]
  total_items: number
}

export type Year = number
export type Day = number
export type Month =
  | 'January'
  | 'February'
  | 'March'
  | 'April'
  | 'May'
  | 'June'
  | 'July'
  | 'August'
  | 'September'
  | 'October'
  | 'November'
  | 'December'

export type MonthRecord = Record<number, Month>

export type Birthdate = `${Year}-${keyof MonthRecord}-${Day}`
