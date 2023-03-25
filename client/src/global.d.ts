export {}

declare global {
  type Medicine = {
    id: string
    description: string
    quantity: number
    unit: string
    batch: string
    expiry: string
    price: number
    createdAt: string
    updatedAt: string
  }
}
