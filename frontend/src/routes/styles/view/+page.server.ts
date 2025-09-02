import { getBuildings } from '$lib/api';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params }) => {
    return {
        observations: await getBuildings()

    };
};