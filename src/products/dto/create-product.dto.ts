export class CreateProductDto {
  id: number;
  title: string;
  description: string;
  productKey: string;
  imageUrl: string;
  rating: number;
  price: number;
  currency: string;
  category: string;
  website: string;
}
