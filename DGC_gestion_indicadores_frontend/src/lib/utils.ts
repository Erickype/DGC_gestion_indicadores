import { type ClassValue, clsx } from "clsx";
import { twMerge } from "tailwind-merge";
import { cubicOut } from "svelte/easing";
import type { TransitionConfig } from "svelte/transition";
import type { CommonError } from "./api/model/errors";
import { toast } from "svelte-sonner";
import { message, type ErrorStatus, type SuperValidated } from "sveltekit-superforms";

export function cn(...inputs: ClassValue[]) {
	return twMerge(clsx(inputs));
}

type FlyAndScaleParams = {
	y?: number;
	x?: number;
	start?: number;
	duration?: number;
};

export const flyAndScale = (
	node: Element,
	params: FlyAndScaleParams = { y: -8, x: 0, start: 0.95, duration: 150 }
): TransitionConfig => {
	const style = getComputedStyle(node);
	const transform = style.transform === "none" ? "" : style.transform;

	const scaleConversion = (
		valueA: number,
		scaleA: [number, number],
		scaleB: [number, number]
	) => {
		const [minA, maxA] = scaleA;
		const [minB, maxB] = scaleB;

		const percentage = (valueA - minA) / (maxA - minA);
		const valueB = percentage * (maxB - minB) + minB;

		return valueB;
	};

	const styleToString = (
		style: Record<string, number | string | undefined>
	): string => {
		return Object.keys(style).reduce((str, key) => {
			if (style[key] === undefined) return str;
			return str + `${key}:${style[key]};`;
		}, "");
	};

	return {
		duration: params.duration ?? 200,
		delay: 0,
		css: (t) => {
			const y = scaleConversion(t, [0, 1], [params.y ?? 5, 0]);
			const x = scaleConversion(t, [0, 1], [params.x ?? 0, 0]);
			const scale = scaleConversion(t, [0, 1], [params.start ?? 0.95, 1]);

			return styleToString({
				transform: `${transform} translate3d(${x}px, ${y}px, 0) scale(${scale})`,
				opacity: t
			});
		},
		easing: cubicOut
	};
};

// Toast errors shown from message of superforms
export function manageToastFromInvalidAddForm() {
	const error = "Formulario inválido"
	const detail = "llenar todos los campos!"
	return toast.error(`${error}: ${detail}`);
}

export function manageToastFromErrorMessageOnAddForm(message: App.Superforms.Message) {
	if (message.error as CommonError) {
		const commonError = message!.error as CommonError;
		return toast.error(`${commonError.message}: ${commonError.detail}`);
	} else {
		const error = message.error as string;
		return toast.error(error);
	}
}

// Superfrom message error generation
export function generateFormMessageFromHttpResponse(form: SuperValidated<Record<string, unknown>>, response: unknown) {
	if ((response as CommonError).status) {
		const error = response as CommonError
		const status = error.status as unknown as ErrorStatus
		return message(form,
			{ success: false, error: error },
			{ status: status })
	}
	return message(form,
		{ success: true, data: response },
	)
}

export function generateFormMessageFromInvalidForm(form: SuperValidated<Record<string, unknown>>) {
	return message(form,
		{ success: false, error: "Invalid form" },
		{ status: 400 })
}

// Common error type generation
export function generateCommonErrorFromFetchError(error: unknown | any): CommonError {
	const errorMessage: CommonError = {
		status: "500",
		status_code: 500,
		detail: (error as Error).message,
		message: "Error al solicitar datos"
	}
	return errorMessage
}


// DraggableFilterStructure
export interface DragableFilterFields {
	id: number;
	key: string;
	name: string;
	label: string;
}