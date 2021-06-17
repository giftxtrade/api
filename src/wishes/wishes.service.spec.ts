import { Test, TestingModule } from '@nestjs/testing';
import { WishesService } from './wishes.service';

describe('WishesService', () => {
  let service: WishesService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [WishesService],
    }).compile();

    service = module.get<WishesService>(WishesService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });
});
