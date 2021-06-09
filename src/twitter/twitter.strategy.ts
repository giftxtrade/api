import { PassportStrategy } from '@nestjs/passport';
import { Strategy, VerifyCallback } from 'passport-twitter';
import { TWITTER } from '../../auth-tokens.json'
import { frontend } from 'src/util/external-routes'

import { Injectable } from '@nestjs/common';

@Injectable()
export class TwitterStrategy extends PassportStrategy(Strategy, 'twitter') {

  constructor() {
    super({
      consumerKey: TWITTER.API_KEY,
      consumerSecret: TWITTER.API_SECRET_KEY,
      callbackURL: frontend.twitter,
      passReqToCallback: true,
      includeEmail: true,
      skipExtendedUserProfile: false,
    });
  }

  async validate(req: Request, accessToken: string, refreshToken: string, profile: any, done: VerifyCallback) {
    const { name, emails, photos } = profile
    const user = {
      email: emails[0].value,
      firstName: name.givenName,
      lastName: name.familyName,
      picture: photos[0].value,
      accessToken
    }
    done(null, user);
  }
}