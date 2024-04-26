import 'svelte';
import type { Locals as DefaultLocals } from '@sveltejs/kit';

declare global {
  namespace App {
    interface Locals extends DefaultLocals {
      globalData: { access_token: string }; // Define the structure of your custom data
    }
  }
}