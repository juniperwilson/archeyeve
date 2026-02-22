import { fetchStyle, doSearch } from '$lib/api';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params }) => {
	return {
		style: await fetchStyle(params.stylename),
		observations: await doSearch({})
	};
};