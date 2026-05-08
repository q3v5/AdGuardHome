import { handleActions } from 'redux-actions';

import * as actions from '../actions/dnsConfig';
import { ALL_INTERFACES_IP, BLOCKING_MODES, DNS_REQUEST_OPTIONS } from '../helpers/constants';

export const DEFAULT_BLOCKING_IPV4 = ALL_INTERFACES_IP;
export const DEFAULT_BLOCKING_IPV6 = '::';

const dnsConfig = handleActions(
    {
        [actions.getDnsConfigRequest.toString()]: (state: any) => ({
            ...state,
            processingGetConfig: true,
        }),
        [actions.getDnsConfigFailure.toString()]: (state: any) => ({
            ...state,
            processingGetConfig: false,
        }),
        [actions.getDnsConfigSuccess.toString()]: (state: any, { payload }: any) => {
            const {
                blocking_ipv4,
                blocking_ipv6,
                upstream_dns,
                upstream_mode,
                fallback_dns,
                bootstrap_dns,
                local_ptr_upstreams,
                ratelimit_whitelist,
                dns_request_device,
                ...values
            } = payload;

            return {
                ...state,
                ...values,
                blocking_ipv4: blocking_ipv4 || DEFAULT_BLOCKING_IPV4,
                blocking_ipv6: blocking_ipv6 || DEFAULT_BLOCKING_IPV6,
                upstream_dns: (upstream_dns && upstream_dns.join('\n')) || '',
                fallback_dns: (fallback_dns && fallback_dns.join('\n')) || '',
                bootstrap_dns: (bootstrap_dns && bootstrap_dns.join('\n')) || '',
                local_ptr_upstreams: (local_ptr_upstreams && local_ptr_upstreams.join('\n')) || '',
                ratelimit_whitelist: (ratelimit_whitelist && ratelimit_whitelist.join('\n')) || '',
                dns_request_device: {
                    enabled: dns_request_device?.enabled ?? false,
                    user_agent: dns_request_device?.user_agent ?? '',
                },
                processingGetConfig: false,
                upstream_mode: upstream_mode === '' ? DNS_REQUEST_OPTIONS.LOAD_BALANCING : upstream_mode,
            };
        },

        [actions.setDnsConfigRequest.toString()]: (state: any) => ({
            ...state,
            processingSetConfig: true,
        }),
        [actions.setDnsConfigFailure.toString()]: (state: any) => ({
            ...state,
            processingSetConfig: false,
        }),
        [actions.setDnsConfigSuccess.toString()]: (state, { payload }: any) => {
            const {
                dns_request_device_enabled,
                dns_request_device_user_agent,
                ...rest
            } = payload;

            return {
                ...state,
                ...rest,
                dns_request_device: {
                    enabled: !!dns_request_device_enabled,
                    user_agent: dns_request_device_user_agent || '',
                },
                processingSetConfig: false,
            };
        },
    },
    {
        processingGetConfig: false,
        processingSetConfig: false,
        blocking_mode: BLOCKING_MODES.default,
        ratelimit: 20,
        blocking_ipv4: DEFAULT_BLOCKING_IPV4,
        blocking_ipv6: DEFAULT_BLOCKING_IPV6,
        blocked_response_ttl: 10,
        upstream_timeout: 10,
        edns_cs_enabled: false,
        disable_ipv6: false,
        dnssec_enabled: false,
        upstream_dns_file: '',
        client_id: '',
        dns_request_device: {
            enabled: false,
            user_agent: '',
        },
    },
);

export default dnsConfig;
