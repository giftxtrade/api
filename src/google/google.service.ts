import { Injectable } from '@nestjs/common';
import { AuthService } from 'src/auth/auth.service';

@Injectable()
export class GoogleService {
  constructor(private readonly authServices: AuthService) { }

  googleLogin(user: {
    accessToken: string,
    email: string,
    firstName: string,
    lastName: string,
    picture: string
  }) {
    return this.authServices
      .login(
        {
          email: user.email,
          name: `${user.firstName} ${user.lastName}`,
          imageUrl: user.picture
        },
        user.accessToken
      );
  }
}