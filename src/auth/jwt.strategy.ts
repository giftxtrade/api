import { ExtractJwt, Strategy } from 'passport-jwt';
import { PassportStrategy } from '@nestjs/passport';
import { Injectable, UnauthorizedException } from '@nestjs/common';
import { AuthService } from './auth.service';
import { JWT } from '../../auth-tokens.json';
import { User } from 'src/users/entities/user.entity';
import JwtPayload from 'src/util/jwtPayload';

@Injectable()
export class JwtStrategy extends PassportStrategy(Strategy) {
  constructor(private authServices: AuthService) {
    super({
      jwtFromRequest: ExtractJwt.fromAuthHeaderAsBearerToken(),
      ignoreExpiration: false,
      secretOrKey: JWT.SECRET,
    });
  }

  async validate(payload: JwtPayload): Promise<JwtPayload> {
    const user: User = await this.authServices.validateUser(payload.user.email);
    if (user == null)
      throw new UnauthorizedException();
    return payload;
  }
}