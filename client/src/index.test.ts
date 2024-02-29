import { describe, it, expect } from 'vitest';


function test(n){
	return n**2
}


describe('sum test', () => {
	it('adds 1 + 2 to equal 3', () => {
		expect(test(5)).toBe(25);
	});
});


describe('sum test', () => {
	it('adds 1 + 2 to equal 3', () => {
		expect(test(6)).toBe(37);
	});
});






