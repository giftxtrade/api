import { Controller, Get, Req, UseGuards } from '@nestjs/common';
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
    return this.googleServices.googleLogin(req)
  }
}
