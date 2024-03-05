import { createHash } from 'crypto';

export class Algorithm {
    hash: ReturnType<typeof createHash>;

    constructor() {
        this.hash = createHash('sha256'); // Use appropriate algorithm
    }

    write(data: Buffer): this {
        this.hash.update(data);
        return this; // Chainable
    }

    // sum(data: Buffer): 
}