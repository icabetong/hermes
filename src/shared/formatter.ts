export function getCurrencyFormatter(withDecimal = true): Intl.NumberFormat {
  return Intl.NumberFormat('en-PH', {
    style: 'currency',
    currency: 'PHP',
    maximumFractionDigits: withDecimal ? 2 : 0,
    minimumFractionDigits: withDecimal ? 2 : 0
  })
}

export function getDateFormatter(
  dateStyle: 'medium' | 'short' | 'full' | 'long' = 'medium'
): Intl.DateTimeFormat {
  return Intl.DateTimeFormat('en-PH', { dateStyle: dateStyle })
}
