<template>
  <bcs-form class="edit-form" v-if="data">
    <bcs-form-item :label="$t('配置信息')">
      <div class="border border-[#DCDEE5] border-solid py-[10px] pr-[24px]">
        <bcs-form-item
          :label="$t('命名空间')" class="config-form-item">
          {{ namespaces || $t('所有') }}
        </bcs-form-item>
        <bcs-form-item :label="$t('关联标签')" class="config-form-item">
          <div
            class="flex flex-col w-full"
            v-if="data.rule.config.label_selector.match_labels && data.rule.config.label_selector.match_labels.length">
            <LogLabel
              class="w-full"
              v-for="item, index in data.rule.config.label_selector.match_labels"
              :value="item"
              :key="index"
              :deleteable="false" />
          </div>
          <span v-else>--</span>
        </bcs-form-item>
        <bcs-form-item :label="$t('工作负载')" class="config-form-item">
          <template v-if="data.rule.config.container.workload_type">
            <div class="flex items-center h-[22px]">
              <div class="flex items-center h-[22px] bg-[#F0F1F5] rounded-sm px-[8px]">
                {{ $t('应用类型') }}: {{ data.rule.config.container.workload_type }}
              </div>
              <div class="flex items-center h-[22px] bg-[#F0F1F5] rounded-sm px-[8px] ml-[8px]">
                {{ $t('应用名称') }}: {{ data.rule.config.container.workload_name }}
              </div>
            </div>
          </template>
          <span v-else>--</span>
        </bcs-form-item>
        <bcs-form-item :label="$t('容器名')" class="config-form-item">
          <template v-if="data.rule.config.container.container_name">
            <span
              class="flex items-center h-[22px] bg-[#F0F1F5] rounded-sm px-[8px] mr-[8px]"
              v-for="item in data.rule.config.container.container_name.split(',')" :key="item">
              {{ item }}
            </span>
          </template>
          <span v-else>--</span>
        </bcs-form-item>
        <template v-if="data.rule.config.paths && !!data.rule.config.paths.length">
          <bcs-form-item :label="$t('日志路径')" class="config-form-item mb-[4px]">
            <div class="mt-[2px]">
              <div
                v-for="item, index in data.rule.config.paths"
                :key="index"
                class="flex items-center leading-none mt-[8px]">
                {{ item }}
              </div>
              <bcs-button
                text
                class="!text-[12px]"
                :disabled="data.status === 'TERMINATED'"
                @click="openLink(data.entrypoint && data.entrypoint.file_log_url)">
                {{ $t('查询文件日志') }}
              </bcs-button>
            </div>
          </bcs-form-item>
          <bcs-form-item :label="$t('日志字符集')" class="config-form-item">
            {{ data.rule.config.data_encoding || '--' }}
          </bcs-form-item>
        </template>
        <bcs-form-item :label="$t('标准输出')" class="config-form-item" v-if="data.rule.config.enable_stdout">
          <bcs-button
            text
            class="!text-[12px] h-[32px] flex items-center"
            :disabled="data.status === 'TERMINATED'"
            @click="openLink(data.entrypoint && data.entrypoint.std_log_url)">
            <div class="h-[32px] relative top-[1px]">{{ $t('查询标准输出日志') }}</div>
          </bcs-button>
        </bcs-form-item>
        <bcs-form-item :label="$t('过滤内容')" class="config-form-item">
          <!-- 字符串过滤 -->
          <span
            class="flex items-center"
            v-if="data.rule.config.conditions.type === 'match' && data.rule.config.conditions.match_content">
            <span class="flex items-center h-[22px] bg-[#F0F1F5] rounded-sm px-[8px] mr-[8px]">
              {{ $t('字符串过滤') }}
            </span>
            <span class="flex items-center h-[22px] bg-[#F0F1F5] rounded-sm px-[8px] mr-[8px]">
              {{data.rule.config.conditions.match_type }}
            </span>
            <span
              class="bcs-ellipsis flex-1 leading-[22px] h-[22px] bg-[#F0F1F5] rounded-sm px-[8px] mr-[8px]">
              {{ data.rule.config.conditions.match_content }}
            </span>
          </span>
          <!-- 分隔符过滤 -->
          <span
            class="flex flex-col mt-[6px]"
            v-else-if="data.rule.config.conditions.type === 'separator'
              && data.rule.config.conditions.separator_filters.length">
            <span class="flex items-center mb-[8px]">
              <span class="flex items-center h-[22px] bg-[#F0F1F5] rounded-sm px-[8px] mr-[8px]">
                {{ $t('分隔符过滤') }}
              </span>
              <span class="flex items-center h-[22px] bg-[#F0F1F5] rounded-sm px-[8px] mr-[8px]">
                {{ data.rule.config.conditions.separator }}
              </span>
              <span class="flex items-center h-[22px] bg-[#F0F1F5] rounded-sm px-[8px] mr-[8px]">
                {{ $t('条件关系') }}
              </span>
              <span class="flex items-center h-[22px] bg-[#F0F1F5] rounded-sm px-[8px] mr-[8px]">
                {{ opMap[filtersOp] }}
              </span>
            </span>
            <span
              class="flex items-center mb-[8px]"
              v-for="item, index in data.rule.config.conditions.separator_filters" :key="index">
              <span class="flex items-center h-[22px] bg-[#F0F1F5] rounded-sm px-[8px] mr-[8px]">
                {{ $t('第 {0} 列', [item.fieldindex]) }}
                <span class="text-[#FF9C01] px-[8px]">{{ item.op }}</span>
                <span class="bcs-ellipsis flex-1">{{ item.word }}</span>
              </span>
              <!-- <span class="flex items-center h-[22px] bg-[#F0F1F5] rounded-sm px-[8px] mr-[8px]"
                    v-if="index <= (data.rule.config.conditions.separator_filters.length - 2)">
                {{opMap[item.logic_op] }}
              </span> -->
            </span>
          </span>
          <span v-else>--</span>
        </bcs-form-item>
      </div>
    </bcs-form-item>
    <bcs-form-item :label="$t('附加日志标签')" class="!mt-[16px]">
      <div class="flex flex-col w-full">
        <div class="flex flex-col w-full" v-if="data.rule.extra_labels && data.rule.extra_labels.length">
          <LogLabel
            :class="[
              'w-full',
              {
                '!mb-0': index === (data.rule.extra_labels.length - 1)
              }
            ]"
            v-for="item, index in data.rule.extra_labels"
            :value="item"
            :key="index"
            :deleteable="false" />
        </div>
        <div v-else>--</div>
        <bcs-checkbox class="mt-[12px]" disabled :value="data.rule.add_pod_label">
          {{ $t('自动添加 Pod 中的 labels') }}
        </bcs-checkbox>
      </div>
    </bcs-form-item>
    <bcs-form-item :label="$t('最近更新人')">{{ data.updator }}</bcs-form-item>
    <bcs-form-item :label="$t('最近更新时间')">{{ data.updated_at }}</bcs-form-item>
    <bcs-form-item :label="$t('备注')">{{ data.description || '--' }}</bcs-form-item>
  </bcs-form>
</template>
<script setup lang="ts">
import { PropType, computed, ref } from 'vue';
import { IRuleData } from './use-log';
import LogLabel from './log-label.vue';
import $i18n from '@/i18n/i18n-setup';

const props = defineProps({
  data: {
    type: Object as PropType<IRuleData>,
    default: () => null,
  },
});

const namespaces = computed(() => props.data.rule?.config?.namespaces?.join(';'));

// 打开日志链接
const openLink = (link) => {
  if (!link) return;
  window.open(link);
};

const filtersOp = computed(() => props.data.rule?.config?.conditions?.separator_filters?.[0]?.logic_op);

const opMap = ref({
  and: $i18n.t('并'),
  or: $i18n.t('或'),
});
</script>
<style lang="postcss" scoped>
.edit-form {
  >>> .bk-label-text {
    font-size: 12px !important;
  }
  >>> .bk-form-content {
    width: 420px;
    display: flex;
    align-items: center;
  }
  >>> .config-form-item {
    margin-top: 0px !important;
    &::before {
      display: none;
    }
    .bk-label {
      color: #979BA5;
    }
  }
  >>> .bk-form-item+.bk-form-item:not(.config-form-item) {
    margin-top: 8px !important;
  }
}
</style>
