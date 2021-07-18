import { Injectable } from '@nestjs/common';
import { AuthService } from 'src/auth/auth.service';

@Injectable()
export class GoogleService {
  constructor(private readonly authServices: AuthService) { }

  async googleLogin(user: {
    accessToken: string,
    email: string,
    firstName: string,
    lastName: string,
    picture: string
  }) {
    const fullname = user.lastName ? `${user.firstName} ${user.lastName}` : user.firstName;
    return await this.authServices
      .login(
        {
          email: user.email,
          name: fullname,
          imageUrl: user.picture
        },
        user.accessToken
      );
  }
}