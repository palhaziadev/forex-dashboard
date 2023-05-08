export interface ForexData {
  CurrencyPair: string
  Value: number
}

export interface CandlestickData {
  open: number
  close: number
  high: number
  low: number
  x: number
}