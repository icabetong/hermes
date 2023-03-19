export {}

declare global {
  type Medicine = {
    id: string
    description: string
    quantity: number
    unit: string
    batch: string
    expiry: number
    price: number
    createdAt: number
    updatedAt: number
  }
}
