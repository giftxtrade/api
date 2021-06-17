import { Test, TestingModule } from '@nestjs/testing';
import { WishesController } from './wishes.controller';
import { WishesService } from './wishes.service';

describe('WishesController', () => {
  let controller: WishesController;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [WishesController],
      providers: [WishesService],
    }).compile();

    controller = module.get<WishesController>(WishesController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });
});
