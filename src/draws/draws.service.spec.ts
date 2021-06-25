import { Test, TestingModule } from '@nestjs/testing';
import { DrawsService } from './draws.service';

describe('DrawsService', () => {
  let service: DrawsService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [DrawsService],
    }).compile();

    service = module.get<DrawsService>(DrawsService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });
});
