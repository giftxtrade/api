import { User } from "src/users/entities/user.entity";

export default interface JwtPayload {
  user: User;
  gToken: string;
}

export interface JwtAuthReturn {
  loggedIn: boolean;
  user: User;
  accessToken: string;
  gToken: string;
}