import { getObservation } from '$lib/api';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params }) => {
	return {
		observation: await getObservation(parseInt(params.obsid))
	};
};