import { Controller, Get, HttpException, HttpStatus, Req, UseGuards } from '@nestjs/common';
import { GoogleService } from 'src/google/google.service';
import { AuthGuard } from '@nestjs/passport';

@Controller('auth/google')
export class GoogleController {
  constructor(private readonly googleServices: GoogleService) { }

  @Get()
  @UseGuards(AuthGuard('google'))
  async googleAuth(@Req() req) { }

  @Get('redirect')
  @UseGuards(AuthGuard('google'))
  googleAuthRedirect(@Req() req) {
    const user = req.user;

    if (!user)
      throw new HttpException({
        message: 'Something went wrong while trying to authenticate'
      }, HttpStatus.BAD_REQUEST)

    return this.googleServices.googleLogin({
      accessToken: user.accessToken,
      email: user.email,
      firstName: user.firstName,
      lastName: user.lastName,
      picture: user.picture
    })
  }
}
