import { PUBLIC_LOCAL_PATH } from "$env/static/public";

			export async function imageStore(formData,id:string,type:string = "default") {
				const fetchResp = await fetch(
					PUBLIC_LOCAL_PATH + `/api/v1/auth/images/${id}/${type}`,
					{
						method: 'POST',
						headers: {
							'Access-Control-Request-Method': 'POST'
						},
						credentials: 'include',
						body: formData
					}
				);
				const json = await fetchResp.json();
                return json
			}