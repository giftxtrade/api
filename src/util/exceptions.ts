import { HttpException, HttpStatus } from '@nestjs/common';

export function BAD_REQUEST(message: string) {
  return new HttpException({
    message: message
  }, HttpStatus.BAD_REQUEST);
}

export function NOT_FOUND(message: string) {
  return new HttpException({
    message: message
  }, HttpStatus.NOT_FOUND);
}

export function UNAUTHORIZED(message: string) {
  return new HttpException({
    message: message
  }, HttpStatus.UNAUTHORIZED);
}