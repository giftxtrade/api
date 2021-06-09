import { Injectable } from '@nestjs/common';
import { User } from 'src/users/entities/user.entity';
import { UsersService } from 'src/users/users.service';
import { JwtService } from '@nestjs/jwt';
import { CreateUserDto } from '../users/dto/create-user.dto';

@Injectable()
export class AuthService {
  constructor(
    private readonly userServices: UsersService,
    private readonly jwtService: JwtService
  ) { }

  async validateUser(email: string, pass: string): Promise<User> {
    return this.userServices.findOne(email);
  }

  async login(createUser: CreateUserDto, accessToken: string): Promise<{ user: User, accessToken: string }> {
    const user = await this.userServices.findOneOrCreate(createUser);
    const payload = { user, accessToken };
    return {
      user: user,
      accessToken: this.jwtService.sign(payload),
    };
  }
}
