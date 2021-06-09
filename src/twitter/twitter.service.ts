import { Injectable } from '@nestjs/common';

@Injectable()
export class TwitterService {
  twitterLogin(req) {
    if (!req.user) {
      return 'No user from Twitter'
    }

    return {
      message: 'User information from google',
      user: req.user
    }
  }
}
