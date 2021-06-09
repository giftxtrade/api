export class CreateUserDto {
  id: number;
  name: string;
  email: string;
  imageUrl: string;
  phone: string = '';
}
