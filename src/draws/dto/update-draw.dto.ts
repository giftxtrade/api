import { PartialType } from '@nestjs/mapped-types';
import { CreateDrawDto } from './create-draw.dto';

export class UpdateDrawDto extends PartialType(CreateDrawDto) {}
