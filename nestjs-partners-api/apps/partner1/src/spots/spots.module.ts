import { Module } from '@nestjs/common';
import { SpotsCoreModule } from '@app/core/spots/spots-core.module';

@Module({
    imports: [SpotsCoreModule],
})
export class SpotsModule { }
