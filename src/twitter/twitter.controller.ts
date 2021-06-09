import { Controller, Get, Req, UnauthorizedException, UseGuards } from '@nestjs/common';
import { AuthGuard } from '@nestjs/passport';
import { TwitterService } from './twitter.service';

@Controller('auth/twitter')
export class TwitterController {
  constructor(private readonly twitterServices: TwitterService) { }

  @Get()
  @UseGuards(AuthGuard('twitter'))
  async twitterAuth(@Req() res) {
    throw new UnauthorizedException();
  }

  @Get('redirect')
  @UseGuards(AuthGuard('twitter'))
  twitterAuthRedirect(@Req() req) {
    return this.twitterServices.twitterLogin(req)
  }
}
